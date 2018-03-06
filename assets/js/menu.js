function getCookie(name) {
    var matches = document.cookie.match(new RegExp(
      "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : undefined;
  }

function checkCookie() {
        var userID = getCookie("auth_session");
        if (userID.length>144) {
            var li1 = document.createElement('li');
            li1.className = "nav-item";
            li1.innerHTML = "<a class=\"nav-link\" href=\"/signOut\">Вихід</a>";
            menu.appendChild(li1);

            var li2 = document.createElement('li');
            li2.className = "nav-item";
            li2.innerHTML = "<a class=\"nav-link\" href=\"/cabinet\">Кабінет</a>";
            menu.appendChild(li2);
        } else {
            var li1 = document.createElement('li');
            li1.className = "nav-item";
            li1.innerHTML = "<a class=\"nav-link\" href=\"/signIn\">Вхід</a>";
            menu.appendChild(li1);

            var li2 = document.createElement('li');
            li2.className = "nav-item";
            li2.innerHTML = "<a class=\"nav-link\" href=\"/signUp\">Реєстрація</a>";
            menu.appendChild(li2);
        }
    }