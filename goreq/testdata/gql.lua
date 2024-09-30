post({
  url = "https://example.com/graphql",
})

headers({
  ["Content-Type"] = "application/json",
})

gql({
  query = [[
    query Continents($code: String!) {
      continents(filter: {code: {eq: $code}}) {
        code
        name
      }
    }
  ]],
  variables = {
    code = "EU",
  },
})
