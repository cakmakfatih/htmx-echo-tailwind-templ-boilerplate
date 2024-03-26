package controllers

type Controller interface {
	register()
}

func RegisterController(c Controller) {
	c.register()
}
