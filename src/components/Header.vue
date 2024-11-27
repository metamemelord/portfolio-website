<template>
  <div class="page-header">
    <div v-show="true">
      <b :style="{'color': currentColor}">New!</b> Read tech blogs <a href="https://medium.gaurav.dev" target="blank">here!</a>
    </div>
    <header class="main-header" :class="{'main-header__shadow': isAtTop}">
      <div class="main-header__wrapper">
        <div class="main-header__nav">
          <div class="main-header__item main-header__home">
            <router-link @click.native="contactMeDialogFunc(false)" to="/" :class="{'inactive':contactMeDialog}" active-class="active" exact>
              <i class="fa fa-home" style="font-size:1.5rem"></i>
              <span class="main-header__item-text">Home</span>
            </router-link>
          </div>
          <div class="main-header__item main-header__blogs">
            <router-link @click.native="contactMeDialogFunc(false)" to="/blogs" :class="{'inactive':contactMeDialog}" active-class="active">
              <i class="fa-solid fa-blog" style="font-size:1.5rem"></i>
              <span class="main-header__item-text">Blogs</span>
            </router-link>
          </div>
        </div>
        <div class="main-header__item main-header__email">
          <span @click="contactMeDialogFunc(true)" :class="{'active':contactMeDialog}">
            <i class="fa-regular fa-envelope" style="font-size:1.5rem"></i>
            <span class="main-header__item-text">Contact me!</span>
          </span>
        </div>
      </div>
    </header>
  </div>
</template>

<script>

export default {
  data() {
    return { position: 0, currentColor: '#F66', colorsEnum: ['#F66', '#87FD05', '#FE02A2', '#FFF', '#FF3503', '#00A0A0', '#FB9214, #07D5E6'] };
  },
  created() {
    let vm = this;
    window.addEventListener("scroll", function() {
      vm.position = this.scrollY;
    });
  },
  computed: {
    isAtTop() {
      return this.position > 20;
    },
    contactMeDialog() {
      return this.$store.state.contactMeDialog;
    }
  }, 
  methods: {
    contactMeDialogFunc(open) {
      this.$store.dispatch('setContactMeDialog', open);
    }
  },
  watch: {
    position() {
      let vm = this;
      setTimeout(() => {
        let col = Math.floor(2 * (Math.random())*vm.colorsEnum.length) % vm.colorsEnum.length;
        if (vm.position != 0) {
          vm.currentColor = vm.colorsEnum[col];
        } else {
          vm.currentColor = vm.colorsEnum[0];
        }
      }, 250);
    }
  }
};
</script>
<style>
.page-header {
  position: fixed;
  z-index: 100000;
  top: 0;
  min-width: 25rem;
  width: 100%;
}

.page-header > div {
  padding: 0.4rem;
  background: #000;
  text-align: center;
  color: #fff;
}

.page-header > div a {
  text-decoration: none;
  cursor: pointer;
  color: var(--accent-color);
}

.page-header > div b {
  transition: 250ms color;
}

.main-header {
  color: white;
  background: #333;
  transition: 0.4s box-shadow ease-out;
}

.main-header__shadow {
  box-shadow: 0px 4px 4px var(--shadow-color);
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

.main-header__item a,
.main-header__item > span {
  transition: 0.5s all;
  color: white;
  text-decoration: none;
  cursor: pointer;
}

.main-header__item a:hover,
.main-header__item > span:hover {
  color: var(--accent-color);
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
  
  .main-header__item a,
  .main-header__item > span {
    display: flex;
    align-items: center;
  }

  .main-header__item a i,
  .main-header__item span i {
    padding-left: 0.3rem;
    padding-right: 0.3rem;
  }

  .main-header__item a span,
  .main-header__item span span {
    padding-left: 0.3rem;
    padding-right: 0.3rem;
  }
  .main-header__item-text {
    display: flex;
  }
}
</style>
