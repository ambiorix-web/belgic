#' Home
#' 
#' Render the homepage.
#' 
#' @inheritParams handler
#' 
#' @name views
#' 
#' @keywords internal
home_get <- \(req, res){
  res$render(
    template_path("home.html"),
    list(
      title = "Hello from R", 
      subtitle = "This is rendered with {glue}"
    )
  )
}
