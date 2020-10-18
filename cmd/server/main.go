package main


//Create Server object and start listener
fun main() {

	const addr := "0.0.0.0:8088"

	server := http.Server {
		Addr: addr,
	}

}