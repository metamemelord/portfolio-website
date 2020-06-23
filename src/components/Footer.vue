<template>
  <footer class="page-footer">
    <span class="page-footer__item page-footer__text-technologies">
      Made with
      <font color="red">❤</font> using Go and Vue.js
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
  position: sticky;
  width: 100%;
  bottom: 0;
  display: flex;
  flex-flow: column;
  background: #000;
  justify-content: center;
  align-items: center;
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
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
