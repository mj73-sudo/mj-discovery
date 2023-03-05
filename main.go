package main

func init() {
	LoadEnv()
	LoadLogConfig()
}

func main() {
	r := GetGin()

	r.Run(":" + ApplicationPort)
}
