#' Contact Page
#' 
#' A contact page with a form.
#' 
#' @inheritParams handler
#' 
#' @keywords internal
contact_get <- \(req, res) {
  res$render(
    template_path("contact.html")
  )
}

#' @keywords internal
contact_post <- \(req, res) {
  body <- parse_multipart(req)
  print(body)
  res$sendf("Thanks %s! Your message was sent!", body$name)
}
