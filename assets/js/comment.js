function getCookie(name) {
    var matches = document.cookie.match(new RegExp(
      "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : "";
  }

function comentCheck() {
        var userID = getCookie("auth_session");
       
        if (userID.length>144) {
            document.getElementById("comment_form").innerHTML = " <div class=\"form-control\"><label>Comment</label><textarea name=\"comment_body\" style=\"width: 100%; height: 60px;\"></textarea><input type=\"hidden\" name=\"article_id\" value=\"{{.Article.ArticleID}}\"></div><button type=\"submit\">Post</button> "
            } else {
                document.getElementById("comment_form").style.visibility = ""
            } 
}