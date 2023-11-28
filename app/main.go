package main

import (
	"github.com/labstack/echo/v4"
	_articleHtpp "luthfi/pemilu/article/delivery/http"
	_articleRepository "luthfi/pemilu/article/repository"
	_articleUseCase "luthfi/pemilu/article/usecase"
	"luthfi/pemilu/internal/components"
	"luthfi/pemilu/internal/config"
	_partaiHandler "luthfi/pemilu/partai/delivery/http"
	_partaiRepository "luthfi/pemilu/partai/repository"
	_partaiUseCase "luthfi/pemilu/partai/usecase"
	_paslonHandler "luthfi/pemilu/paslon/delivery/http"
	_paslonRepository "luthfi/pemilu/paslon/repository"
	_paslonUseCase "luthfi/pemilu/paslon/usecase"
	_userHttp "luthfi/pemilu/user/delivery/http"
	_userRepository "luthfi/pemilu/user/repository"
	_userUseCase "luthfi/pemilu/user/usecase"
)

func main() {
	db := components.GetDatabaseConnection(config.Get())
	cnf := config.Get()
	e := echo.New()
	
	//article
	articleRepository := _articleRepository.NewArticleRepository(db)
	articleUseCase := _articleUseCase.NewArticleUseCase(articleRepository)
	_articleHtpp.NewArticleHandler(e, articleUseCase)
	
	//user
	userRepository := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepository)
	_userHttp.NewUserHandler(e, userUseCase)
	
	//partai
	partaiRepository := _partaiRepository.NewPartaiRepository(db)
	partaiUseCase := _partaiUseCase.NewPartaiUseCase(partaiRepository)
	_partaiHandler.NewPartaiHandler(e, partaiUseCase)
	
	//paslon
	paslonRepository := _paslonRepository.NewPaslonRepository(db)
	paslonUseCase := _paslonUseCase.NewPaslonUseCase(paslonRepository)
	_paslonHandler.NewPaslonHandler(e, paslonUseCase)
	
	e.Logger.Fatal(e.Start(cnf.Server.Host + ":" + cnf.Server.Port))
}
