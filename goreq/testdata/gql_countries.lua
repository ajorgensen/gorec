post({
  url = "https://countries.trevorblades.com/",
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
