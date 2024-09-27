post({
  url = "https://example.com/graphql",
})

headers({
  ["Content-Type"] = "application/json",
})

gql(
  [[
query Continents($code: String!) {
    continents(filter: {code: {eq: $code}}) {
      code
      name
    }
}
]],
  {
    code = "EU",
  }
)
