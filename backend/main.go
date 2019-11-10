package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type Carteira struct {
	Adress string  `json:"address"`
	Id     int     `json:"id"`
	Amount float64 `json:"amount"`
}

var l = make(map[string]*Carteira)

var Cart = Carteira{
	Adress: "Principal",
	Id:     -1,
	Amount: 0.0,
}

type Transaction struct {
	addr   string
	amount float64
}

func main() {
	AdicionaCarteira("Lucas")
	AdicionaCarteira("Pedro")

	RealizarTransacao("Lucas", -1)
	Server()

}

func RealizarTransacao(nome string, amount float64) {
	search := Procurar(nome)
	search.sendTransaction(amount)

}

func Procurar(addr string) *Carteira {
	// c := Carteira{}
	carteira, ok := l[addr]
	if !ok {
		fmt.Println("carteira inexistente")
	}

	return carteira
}

func (c *Carteira) sendTransaction(amount float64) int {
	if amount < 0 {
		fmt.Printf("Erro! numero negativo nao permitido\n")
		return 1
	}
	if c == nil {
		fmt.Printf("Erro! carteira Nao encontrada\n")
		return 1
	}

	if c.Amount < amount {
		fmt.Printf("Erro! Saldo insuficiente")
		return 1
	}

	// remove amount da carteira que esta enviando
	c.Amount -= amount
	// cart 'e a carteira global do sistema
	Cart.Amount += amount
	fmt.Printf("Trasnferencia realizada com sucesso")
	return 0
}

func AdicionaCarteira(nome string) {
	cat := &Carteira{}
	cat.Adress = nome
	cat.Amount = 0.0

	l[nome] = cat
}

func GetLista() {
	return
}

func Server() {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Route("/carteiras", func(r chi.Router) {
		r.Get("/list", ListCarteiras)
		r.Get("/add/nome={nome}", CriarCarteira)

	})

	http.ListenAndServe(":3000", r)
}

func CriarCarteira(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aaaaa")
	nome := chi.URLParam(r, "nome")

	AdicionaCarteira(nome)
}

func ListCarteiras(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, NewCarteirasListResponse()); err != nil {
		render.Render(w, r, nil)
		return
	}
}

func NewCarteirasListResponse() []render.Renderer {
	list := []render.Renderer{}
	for _, cat := range l {
		list = append(list, cat)
	}
	fmt.Println(len(list))
	return list
}

func (c Carteira) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
