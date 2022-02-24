pkg_file <- function(path) {
  system.file(path, package = "/home/john/Golang/src/github.com/devOpifex/belgic/app")
}

#' @export 
assets_path <- function() {
  pkg_file("assets")
}

template_path <- function(...) {
  assets <- pkg_file("templates")
  file.path(assets, ...)
}
