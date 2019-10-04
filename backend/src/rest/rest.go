package rest

import "github.com/gin-gonic/gin"

func RunAPI(address string) error {
	r := gin.Default()

	// get products
	r.GET("/products", func(c *gin.Context) {
		// return a list of all products to the client
	})

	// get promos
	r.GET("/promos", func(c *gin.Context) {
		// return a list of all promotions to the client
	})

	//post user sign in
	r.POST("/users/signin", func(c *gin.Context) {
		//sign in a user
	})
	//add user
	r.POST("/users", func(c *gin.Context) {
		//add a user
	})
	//post user sign out
	/*
	  In the path below, our relative url needs to include the user id
	  Since the id will differ based on the user, the Gin framework allows us to include a wildcard. In Gin, the wildcard will take the form ':id' to indicate that we are expecting a parameter here with the name 'id'
	*/
	r.POST("/user/:id/signout", func(c *gin.Context) {
		//sign out a user with the provided id
	})
	//get user orders
	r.GET("/user/:id/orders", func(c *gin.Context) {
		//get all orders belonging to the provided user id
	})
	//post purchase charge
	r.POST("/users/charge", func(c *gin.Context) {
		//charge credit card for user
	})
	return r.Run(address)
}
