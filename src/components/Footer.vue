<template>
  <footer class="page-footer">
    <span class="page-footer__item page-footer__text-technologies">
      Made with
      <font color="red">❤</font> using Vue.js
    </span>
    <span class="page-footer__item page-footer__text-seperator">|</span>
    <span class="page-footer__item page-footer__text-copyright">© {{year}} Gaurav Saini</span>
  </footer>
</template>

<script>
export default {
  computed: {
    year() {
      return new Date().getFullYear();
    }
  }
};
</script>

<style>
.page-footer {
  width: 100%;
  bottom: 0;
  display: flex;
  flex-flow: column;
  background: #000;
  justify-content: center;
  align-items: center;
}

.page-footer__item {
  color: white;
  padding: 0.3rem;
}

.page-footer__text-seperator {
  display: none;
}

@media screen and (min-width: 550px) {
  .page-footer {
    flex-flow: row;
  }

  .page-footer__text-seperator {
    display: block;
  }
}
</style>
