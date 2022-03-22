#' Build
#' 
#' Build the application
#' 
#' @import ambiorix
#' 
#' @return An object of class `Ambiorix`.
#' 
#' @export 
build <- \() {
  app <- Ambiorix$new()

  # 404 page
  app$not_found <- render_404

  # 500 server errors
  app$error <- render_500

  # serve static files
  app$static(assets_path(), "static")

  # homepage
  app$get("/", home_get)

  # about
  app$get("/about", about_get)

  # contact
  app$get("/contact", contact_get)
  app$post("/contact", contact_post)

  # websocket 
  app$receive("hello", \(msg, ws){
    print(msg)
    ws$send("hello", "Hello back! (sent from R)")
  })

  return(app)
}
