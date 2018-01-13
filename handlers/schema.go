package handlers

import (
	"html/template"
	"net/http"

	graphql "github.com/neelance/graphql-go"
	"github.com/nfisher/graphql-go/relay"
)

const Schema = `
schema {
	query: Query
}

type Query {
	pingStr(v: String!): PongStr!
	pingInt(v: Int!): PongInt!
	# floatEcho(v: Float!): Float!
}

type PongStr {
	value: String!
}

type PongInt {
	value: Int!
}
`

type PongInt int32
type PongStr string

type intResolver struct {
	V int32
}

func (r *intResolver) Value() int32 {
	return r.V
}

type strResolver struct {
	V string
}

func (r *strResolver) Value() string {
	return r.V
}

type EchoResolver struct{}

func (_ EchoResolver) PingStr(args struct{ V string }) *strResolver {
	return &strResolver{args.V}
}

func (_ EchoResolver) PingInt(args struct{ V int32 }) *intResolver {
	return &intResolver{args.V}
}

type Resolver struct {
	EchoResolver
}

var schema = graphql.MustParseSchema(Schema, &Resolver{})
var relayHandler = &relay.Handler{Schema: schema}

var GQLHandler = func(w http.ResponseWriter, req *http.Request) {
	relayHandler.ServeHTTP(w, req)
}

var gqTmpl = template.Must(template.New("ui.html").Parse(GraphiQL))

// GraphiQLHandler handles requests to issue the GraphiQL HTML.
func GraphiQLHandler(w http.ResponseWriter, req *http.Request) {
	page := struct{}{}
	gqTmpl.Execute(w, &page)
}

const GraphiQL = `<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.min.css" integrity="sha256-gSgd+on4bTXigueyd/NSRNAy4cBY42RAVNaXnQDjOW8=" crossorigin="anonymous" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.1/fetch.min.js" integrity="sha256-TQsP3yTWwfvm6Auy90oBeVhYhGZuKa1jRM3vpnQpX+8=" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.6.2/react.min.js" integrity="sha256-c/17te7UpABi7+wcIHAAiIMOrNMVcTIzoxtRTDoYB4s=" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/15.6.2/react-dom.min.js" integrity="sha256-Xhtg7QJuNhwB5AzaUcgr0iqNtCitzN+c/6k5/SOtENU=" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.min.js" integrity="sha256-oeWyQyKKUurcnbFRsfeSgrdOpXXiRYopnPjTVZ+6UmI=" crossorigin="anonymous"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			"use strict";
			function graphQLFetcher(graphQLParams) {
				return fetch("/graphql", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}

			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>`
