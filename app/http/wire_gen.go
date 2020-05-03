// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package http

import (
	"github.com/aristat/golang-example-app/app/auth"
	"github.com/aristat/golang-example-app/app/casbin"
	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/db"
	"github.com/aristat/golang-example-app/app/db/repo"
	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/graphql"
	"github.com/aristat/golang-example-app/app/grpc"
	"github.com/aristat/golang-example-app/app/logger"
	"github.com/aristat/golang-example-app/app/resolver"
	"github.com/aristat/golang-example-app/app/routers/oauth-router"
	"github.com/aristat/golang-example-app/app/routers/products-router"
	"github.com/aristat/golang-example-app/app/routers/users-router"
	"github.com/aristat/golang-example-app/app/session"
	"github.com/aristat/golang-example-app/app/tracing"
)

// Injectors from injector.go:

func Build() (*Http, func(), error) {
	context, cleanup, err := entrypoint.ContextProvider()
	if err != nil {
		return nil, nil, err
	}
	viper, cleanup2, err := config.Provider()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	sessionConfig, cleanup3, err := session.Cfg(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	manager, cleanup4, err := session.Provider(context, sessionConfig)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	loggerConfig, cleanup5, err := logger.ProviderCfg(viper)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	zap, cleanup6, err := logger.Provider(context, loggerConfig)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbConfig, cleanup7, err := db.Cfg(viper)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gormDB, cleanup8, err := db.ProviderGORM(context, zap, dbConfig)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbManager, cleanup9, err := db.Provider(context, zap, dbConfig, gormDB)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	oauth_routerConfig, cleanup10, err := oauth_router.Cfg(viper)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tokenStore, cleanup11, err := oauth_router.TokenStore(oauth_routerConfig)
	if err != nil {
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientStore, cleanup12, err := oauth_router.ClientStore(oauth_routerConfig)
	if err != nil {
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	oauth_routerManager, cleanup13, err := oauth_router.Provider(context, zap, tokenStore, manager, clientStore)
	if err != nil {
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	usersRepo, cleanup14, err := repo.NewUsersRepo(gormDB)
	if err != nil {
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	repoRepo, cleanup15, err := repo.Provider(usersRepo)
	if err != nil {
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	configuration, cleanup16, err := tracing.ProviderCfg(viper)
	if err != nil {
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup17, err := tracing.Provider(context, configuration, zap)
	if err != nil {
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	grpcConfig, cleanup18, err := grpc.Cfg(viper)
	if err != nil {
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	poolManager, cleanup19, err := grpc.Provider(context, tracer, zap, grpcConfig)
	if err != nil {
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceManagers := users_router.ServiceManagers{
		Session:     manager,
		DB:          dbManager,
		Oauth:       oauth_routerManager,
		Repo:        repoRepo,
		PoolManager: poolManager,
	}
	users_routerManager, cleanup20, err := users_router.Provider(context, zap, serviceManagers)
	if err != nil {
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	products_routerServiceManagers := products_router.ServiceManagers{
		PoolManager: poolManager,
	}
	products_routerConfig, cleanup21, err := products_router.Cfg(viper)
	if err != nil {
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	products_routerManager, cleanup22, err := products_router.Provider(context, zap, products_routerServiceManagers, products_routerConfig)
	if err != nil {
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	authConfig, cleanup23, err := auth.ProviderCfg(viper)
	if err != nil {
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	middleware, cleanup24, err := auth.Provider(authConfig, zap)
	if err != nil {
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	resolverConfig, cleanup25, err := resolver.Cfg(viper)
	if err != nil {
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	enforcer, cleanup26, err := casbin.Provider()
	if err != nil {
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	managers := resolver.Managers{
		Repo:        repoRepo,
		PollManager: poolManager,
	}
	graphqlConfig, cleanup27, err := resolver.Provider(context, zap, resolverConfig, enforcer, managers)
	if err != nil {
		cleanup26()
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	config2, cleanup28, err := graphql.Cfg(viper)
	if err != nil {
		cleanup27()
		cleanup26()
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	graphQL, cleanup29, err := graphql.Provider(context, graphqlConfig, zap, config2)
	if err != nil {
		cleanup28()
		cleanup27()
		cleanup26()
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpManagers := Managers{
		session:        manager,
		users:          users_routerManager,
		oauth:          oauth_routerManager,
		products:       products_routerManager,
		authMiddleware: middleware,
		graphql:        graphQL,
	}
	chiMux, cleanup30, err := Mux(httpManagers, zap, tracer)
	if err != nil {
		cleanup29()
		cleanup28()
		cleanup27()
		cleanup26()
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpConfig, cleanup31, err := Cfg(viper)
	if err != nil {
		cleanup30()
		cleanup29()
		cleanup28()
		cleanup27()
		cleanup26()
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	http, cleanup32, err := Provider(context, chiMux, zap, httpConfig, httpManagers)
	if err != nil {
		cleanup31()
		cleanup30()
		cleanup29()
		cleanup28()
		cleanup27()
		cleanup26()
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return http, func() {
		cleanup32()
		cleanup31()
		cleanup30()
		cleanup29()
		cleanup28()
		cleanup27()
		cleanup26()
		cleanup25()
		cleanup24()
		cleanup23()
		cleanup22()
		cleanup21()
		cleanup20()
		cleanup19()
		cleanup18()
		cleanup17()
		cleanup16()
		cleanup15()
		cleanup14()
		cleanup13()
		cleanup12()
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
