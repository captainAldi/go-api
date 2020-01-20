// main.go

package main

func main() {
	a := App{}
	// konfigurasi database disini
	a.Initialize("root", "root", "127.0.0.1", "8889", "db_go")
	a.Run(":8080")
}
