package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/add-to-cart", func(w http.ResponseWriter, r *http.Request) {
		// 获取商品ID和数量
		itemID := r.FormValue("item_id")
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))

		// 获取购物车Cookie
		cookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie {
			// 如果没有购物车Cookie，则创建一个新的购物车Cookie
			cookie = &http.Cookie{
				Name:  "cart",
				Value: "",
			}
		}

		// 将商品添加到购物车
		cart := parseCart(cookie.Value)
		cart[itemID] += quantity

		// 更新购物车Cookie
		cookie.Value = formatCart(cart)
		http.SetCookie(w, cookie)

		fmt.Fprint(w, "Item added to cart.")
	})

	http.HandleFunc("/view-cart", func(w http.ResponseWriter, r *http.Request) {
		// 获取购物车Cookie
		cookie, err := r.Cookie("cart")
		if err == http.ErrNoCookie {
			// 如果没有购物车Cookie，则显示空购物车
			fmt.Fprint(w, "Your cart is empty.")
			return
		}

		// 显示购物车中的商品
		cart := parseCart(cookie.Value)
		if len(cart) == 0 {
			fmt.Fprint(w, "Your cart is empty.")
			return
		}

		fmt.Fprintln(w, "Your cart:")
		for itemID, quantity := range cart {
			fmt.Fprintf(w, "- Item #%s: %d\n", itemID, quantity)
		}
	})

	http.ListenAndServe(":8080", nil)
}

// 将购物车字符串解析为map
func parseCart(cartStr string) map[string]int {
	cart := make(map[string]int)
	for _, item := range split(cartStr, ";") {
		if item == "" {
			continue
		}
		parts := split(item, ":")
		if len(parts) != 2 {
			continue
		}
		id := parts[0]
		quantity, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		cart[id] += quantity
	}
	return cart
}

// 将购物车map格式化为字符串
func formatCart(cart map[string]int) string {
	items := make([]string, 0, len(cart))
	for itemID, quantity := range cart {
		items = append(items, itemID+":"+strconv.Itoa(quantity))
	}
	return join(items, ";")
}

// 将字符串按分隔符分割为切片
func split(s, sep string) []string {
	return strings.Split(s, sep)
}

// 将切片按分隔符拼接为字符串
func join(s []string, sep string) string {
	return strings.Join(s, sep)
}
