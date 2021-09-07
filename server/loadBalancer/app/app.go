package app

func NewApp() App {
	return App{}
}

type App struct {
}

func (this *App) Run() {
	this.starHttpServer()
}
