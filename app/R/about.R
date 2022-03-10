#' About
#' 
#' Render the about.
#' 
#' @inheritParams handler
#' 
#' @name views
#' @keywords internal
about_get <- function(req, res) {
  res$send(
    "About us!"
  )
}
