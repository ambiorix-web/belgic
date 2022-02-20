library(ambiorix)

app <- Ambiorix$new()

app$use(\(req, res){
  req$set(x, 1) # set x to 1
})

app$get("/", \(req, res){
  print(req$get(x)) 
  res$send("Using {ambiorix}!")
})

app$get("/about", \(req, res){
  res$text("About")
})

app$use(\(req, res){
  req$get(x)
})

app$start()
