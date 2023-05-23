package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := PriceDB{}
	db.db = database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

type PriceDB struct {
	sync.Mutex
	db database
}

// 显示所有项
func (p *PriceDB) list(w http.ResponseWriter, req *http.Request) {
	p.Lock()
	for item, price := range p.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	p.Unlock()
}

// 查询
func (p *PriceDB) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p.Lock()
	price, ok := p.db[item]
	p.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q", item)
		return
	}
	fmt.Fprintf(w, "%s", price)
}

// 更新
func (p *PriceDB) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	newprice := req.URL.Query().Get("price")
	p.Lock()
	if price, ok := p.db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q", item)
		return
	} else {
		pf, _ := strconv.ParseFloat(newprice, 32)
		p.db[item] = dollars(pf)
		fmt.Fprintln(w, "Update success!")
		fmt.Fprintf(w, "Old price: %s: %s\n", item, price)
		fmt.Fprintf(w, "New price: %s: %s\n", item, dollars(pf))
	}
	p.Unlock()
}

// 创建
func (p *PriceDB) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	p.Lock()
	if _, ok := p.db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%q is existed", item)
		return
	}
	pf, _ := strconv.ParseFloat(price, 32)
	p.db[item] = dollars(pf)
	fmt.Fprintln(w, "Create success!")
	p.Unlock()
}

// 删除
func (p *PriceDB) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p.Lock()
	if _, ok := p.db[item]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%q is not existed", item)
		return
	}
	delete(p.db, item)
	fmt.Fprintln(w, "Delete success!")
	p.Unlock()
}
