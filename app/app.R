pkgload::load_all(
  export_all = TRUE,
  helpers = FALSE,
  attach_testthat = FALSE
)
library(ambiorix)

app <- Ambiorix$new()

# 404 page
app$not_found <- render_404

# serve static files
app$static(assets_path(), "static")

# homepage
app$get("/", render_home)

# about
app$get("/about", render_about)

# websocket 
app$receive("hello", function(msg, ws){
  print(msg)
  ws$send("hello", "Hello back! (sent from R)")
})

app$start()
