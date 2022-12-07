<header>
  <nav class="navbar navbar-expand-md navbar-dark flex-column flex-md-row bd-navbar bg-dark" id="navbar">
    <a class="navbar-brand mr-0 mr-md-2" href="#">
      <img src="/bmstu.webp" id="navbar-logo" alt="bmstu logo"/>
      CЭЛФ
    </a>


    <div class="navbar-nav-scroll">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item active">
          <a class="nav-link" href="#">Лабораторные работы</a>
        </li>
      </ul>
    </div>

    <!-- TODO: Working toggle -->
    <button class="navbar-toggler collapsed" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>

    <ul id="navbar-login" class="navbar-nav flex-row ml-md-auto d-none d-md-flex">
      <li class="nav-item active">
        <a class="nav-link" href="#">Войти</a>
      </li>
      <li class="nav-item active">
        <a class="nav-link" href="#">Зарегистрироваться</a>
      </li>
    </ul>
  </nav>
</header>

<style>
    #navbar {
        padding: .5rem 1rem;
    }

    #navbar-logo {
        width: 40px;
        height: 40px;
    }

    #navbar-login {
        margin-left: auto !important;
    }
</style>