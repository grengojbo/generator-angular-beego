<ul class="nav navbar-nav navbar-right">
  {{<if .IsLogin>}}
  <li class="dropdown">
    <a href="#" class="dropdown-toggle" data-toggle="dropdown">IMG AVATAR</a>
    <ul class="dropdown-menu">
      {{<if .User.IsAdmin>}}
      <li><a href="{{<.AppUrl>}}admin">{{<i18n .Lang "admin.admin_center">}}</a></li>
      <li class="divider"></li>
      {{<end>}}
      <li><a href="{{<.User.Link>}}"><span>{{<i18n .Lang "user.home">}}</span></a></li>
      <li><a href="{{<.AppUrl>}}settings/profile"><span>{{<i18n .Lang "auth.user_profile">}}</span></a></li>
      <li><a href="{{<.AppUrl>}}logout">{{<i18n .Lang "auth.logout">}}</a></li>
    </ul>
  </li>
  {{<else>}}
  <li class="dropdown{{<if .IsLoginPage>}} active{{<end>}}">
    <a href="/login" class="dropdown-toggle" data-toggle="dropdown">{{<i18n .Lang "auth.login">}}</a>
    <ul class="dropdown-menu" style="padding: 15px;min-width: 250px;">
      <li>
        <form action="/login" method="POST">
          {{<.xsrf_html>}}{{<.once_html>}}
          <div class="alert alert-info alert-small">{{<i18n .Lang "auth.login_your_account">}}</div>
          <div class="form-group">
            <input class="form-control" name="UserName" type="text" placeholder="{{<i18n .Lang "auth.login_username">}} / {{<i18n .Lang "auth.login_email">}}" value="">
          </div>
          <div class="form-group">
            <input class="form-control" name="Password" type="password" placeholder="{{<i18n .Lang "auth.login_password">}}">
          </div>
          <p>
            <label>
              <input type="hidden" name="Remember" value="0">
              <button type="button" data-toggle="button" data-name="Remember" value="0" class="btn btn-default btn-xs btn-checked"><i class="icon icon-ok"></i></button>{{<i18n .Lang "auth.login_remember_me">}}
            </label>
            <span class="pull-right"><a href="/forgot"><i class="icon-question-sign"></i> {{<i18n .Lang "auth.forgot_password">}}</a></span>
          </p>
          <button class="btn btn-primary btn-block" type="submit">{{<i18n .Lang "auth.login">}}</button>
          <div class="auth-socials">
            <a href="login/google" class="icon-google-plus google"></a>
            <a href="/login/github" class="icon-github github"></a>
          </div>
        </form>
      </li>
    </ul>
  </li>
  <li {{<if .IsRegister>}}class="active"{{<end>}}><a href="/register">{{<i18n .Lang "auth.register">}}</a></li>
  {{<end>}}
</ul>