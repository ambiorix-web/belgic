#' About
#' 
#' Render the about.
#' 
#' @inheritParams handler
#' 
#' @name views
#' @keywords internal
about_get <- \(req, res) {
  res$send(
    "About us!"
  )
}
