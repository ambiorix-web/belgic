
#' Error
#' 
#' Rendering errors (!= 200).
#' 
#' @name errors
#' 
#' @keywords internal
render_404 <- function(req, res){
  res$send_file(
    template_path("404.html"),
    status = 404L
  )
}

#' @rdname errors
#' @keywords internal
render_500 <- function(req, res){
  res$send(
    "Internal server error",
    status = 500L
  )
}