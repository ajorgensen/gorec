get({
  url = "https://example.com/",
})

headers({
  ["Content-Type"] = "application/json",
  ["Authorization"] = "Bearer " .. env("API_KEY"),
})
