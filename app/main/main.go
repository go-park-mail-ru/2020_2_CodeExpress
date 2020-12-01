package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

<<<<<<< HEAD:app/main/main.go
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/admin/proto_admin"

	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/track/proto_track"

=======
>>>>>>> 0d81e9f (CP-100: Перенести работу с сессиями на микросервисную архитектуру):app/main.go
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/session/proto_session"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"

	userDelivery "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/user/delivery"
	userRepository "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/user/repository"
	userUsecase "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/user/usecase"

	sessionDelivery "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/session/delivery"
	sessionUsecase "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/session/usecase"

	artistDelivery "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/artist/delivery"
	artistRepository "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/artist/repository"
	artistUsecase "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/artist/usecase"

	trackDelivery "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/track/delivery"
	trackUsecase "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/track/usecase"

	albumDelivery "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/album/delivery"
	albumRepository "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/album/repository"
	albumUsecase "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/album/usecase"

	playlistDelivery "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/playlist/delivery"
	playlistRepository "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/playlist/repository"
	playlistUsecase "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/playlist/usecase"

	searchDelivery "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/search/delivery"
	searchRepository "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/search/repository"
	searchUsecase "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/search/usecase"

	"github.com/go-park-mail-ru/2020_2_CodeExpress/config"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/mwares"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

const (
	argsCount = 2
	argsUsage = "Usage: go run main.go $config_file"
	dbName    = "postgres"
)

func main() {
	if len(os.Args) != argsCount {
		fmt.Println(argsUsage)
		return
	}

	configFileName := os.Args[1]

	conf, err := config.LoadConfig(configFileName)
	if err != nil {
		logrus.Fatal(err)
	}

	dbConn, err := sql.Open(dbName, conf.GetDbConnString())
	if err != nil {
		logrus.Fatal(err)
	}

	defer func() {
		_ = dbConn.Close()
	}()

	if err := dbConn.Ping(); err != nil {
		logrus.Fatal(err)
	}
	log.Printf("DB connected on %s", conf.GetDbConnString())

	userRep := userRepository.NewUserRep(dbConn)
	artistRep := artistRepository.NewArtistRep(dbConn)
	albumRep := albumRepository.NewAlbumRep(dbConn)
	playlistRep := playlistRepository.NewPlaylistRep(dbConn)
	searchRep := searchRepository.NewSearchRep(dbConn)

	userUsecase := userUsecase.NewUserUsecase(userRep)
<<<<<<< HEAD:app/main/main.go
=======
	artistUsecase := artistUsecase.NewArtistUsecase(artistRep)
	trackUsecase := trackUsecase.NewTrackUsecase(trackRep)
>>>>>>> 0d81e9f (CP-100: Перенести работу с сессиями на микросервисную архитектуру):app/main.go
	albumUsecase := albumUsecase.NewAlbumUsecase(albumRep)
	playlistUsecase := playlistUsecase.NewPlaylistUsecase(playlistRep)
	searchUsecase := searchUsecase.NewSearchUsecase(searchRep)

<<<<<<< HEAD:app/main/main.go
	adminGRPCConn, err := grpc.Dial(conf.GetAdminMicroserviceConnString(), grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	defer adminGRPCConn.Close()
	adminGRPC := proto_admin.NewAdminServiceClient(adminGRPCConn)
	artistUsecase := artistUsecase.NewArtistUsecase(artistRep, adminGRPC)

=======
>>>>>>> 0d81e9f (CP-100: Перенести работу с сессиями на микросервисную архитектуру):app/main.go
	sessionGRPCConn, err := grpc.Dial(conf.GetSessionMicroserviceConnString(), grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	defer sessionGRPCConn.Close()
	sessionGRPC := proto_session.NewSessionServiceClient(sessionGRPCConn)
	sessionUsecase := sessionUsecase.NewSessionUsecase(sessionGRPC)

<<<<<<< HEAD:app/main/main.go
	trackGRPCConn, err := grpc.Dial(conf.GetTrackMicroserviceConnString(), grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	defer trackGRPCConn.Close()
	trackGRPC := proto_track.NewTrackServiceClient(trackGRPCConn)
	trackUsecase := trackUsecase.NewTrackUsecase(trackGRPC)

=======
>>>>>>> 0d81e9f (CP-100: Перенести работу с сессиями на микросервисную архитектуру):app/main.go
	userHandler := userDelivery.NewUserHandler(userUsecase, sessionUsecase)
	sessionHandler := sessionDelivery.NewSessionHandler(sessionUsecase, userUsecase)
	artistHandler := artistDelivery.NewArtistHandler(artistUsecase)
	albumHandler := albumDelivery.NewAlbumHandler(albumUsecase, artistUsecase, trackUsecase)
	trackHandler := trackDelivery.NewTrackHandler(trackUsecase, sessionUsecase, userUsecase)
	playlistHandler := playlistDelivery.NewPlaylistHandler(playlistUsecase, trackUsecase)
	searchHandler := searchDelivery.NewSearchHandler(searchUsecase)

	e := echo.New()
	mm := mwares.NewMiddlewareManager(sessionUsecase, userUsecase)

	e.Use(mm.AccessLog, mm.PanicRecovering, mm.CORS(), mm.XSS())
	e.Static("/avatars", "avatars")
	e.Static("/artist_posters", "artist_posters")
	e.Static("/track_audio", "track_audio")
	e.Static("/album_posters", "album_posters")
	e.Static("/artist_avatars", "artist_avatars")
	e.Static("/playlist_posters", "playlist_posters")

	userHandler.Configure(e, mm)
	sessionHandler.Configure(e)
	artistHandler.Configure(e, mm)
	trackHandler.Configure(e, mm)
	albumHandler.Configure(e, mm)
	playlistHandler.Configure(e, mm)
	searchHandler.Configure(e)

	e.Logger.Fatal(e.Start(conf.GetServerConnString()))
}
