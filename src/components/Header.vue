<template>
  <header class="main-header" :class="{'main-header__shadow': isAtTop}">
    <div class="main-header__wrapper">
      <div class="main-header__nav">
        <div class="main-header__item main-header__home">
          <router-link to="/" active-class="active" exact>
            <i class="fa fa-home" style="font-size:1.5rem"></i>
            <span class="main-header__item-text">Home</span>
          </router-link>
        </div>
        <div class="main-header__item main-header__blogs">
          <router-link to="/blogs" active-class="active">
            <i class="fa fa-pencil-square-o" style="font-size:1.5rem"></i>
            <span class="main-header__item-text">Blogs</span>
          </router-link>
        </div>
      </div>
      <div class="main-header__item main-header__email">
        <a href="mailto:hello@gauravsaini.dev" target="blank">
          <i class="fa fa-envelope-o" style="font-size:1.5rem"></i>
          <span class="main-header__item-text">Contact me!</span>
        </a>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  data() {
    return { position: 0 };
  },
  created() {
    let vm = this;
    window.addEventListener("scroll", function(event) {
      vm.position = this.scrollY;
    });
  },
  computed: {
    isAtTop() {
      return this.position > 20;
    }
  }
};
</script>
<style>
.main-header {
  top: 0;
  position: fixed;
  z-index: 100000;
  width: 100%;
  min-width: 20rem;
  color: white;
  background: #333;
  transition: 0.4s box-shadow ease-out;
}

.main-header__shadow {
  box-shadow: 0px 0px 4px 4px rgba(0, 0, 0, 0.35) !important;
}

.main-header__wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  max-width: 70rem;
  margin: 0 auto;
}

.main-header__nav {
  display: flex;
  justify-content: center;
  align-items: center;
}

.main-header__item {
  padding: 1.2rem 2rem;
}

.main-header__item a {
  transition: 0.5s all;
  color: white;
  text-decoration: none;
}

.main-header__item a:hover {
  color: #aaa;
}

.main-header__item-text {
  display: none;
}

@media screen and (min-width: 550px) {
  .main-header__wrapper {
    justify-content: space-between;
  }
  .main-header__item {
    justify-content: center;
    padding: 1.2rem;
    transition: 0.4s all;
  }
  .main-header__item:hover {
    background: #ddd;
    color: #333;
  }
  .main-header__item a {
    color: inherit;
    display: flex;
    align-items: center;
  }

  .main-header__item a:hover {
    color: inherit;
  }
  .main-header__item a i {
    padding-left: 0.3rem;
    padding-right: 0.3rem;
  }

  .main-header__item a span {
    padding-left: 0.3rem;
    padding-right: 0.3rem;
  }
  .main-header__item-text {
    display: flex;
  }
}
</style>
