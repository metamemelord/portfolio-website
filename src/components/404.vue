<template>
  <div class="error__main">
    <div class="parallax-container" id="scene-top">
      <h2 data-depth="0.6">Looks like you are lost</h2>
    </div>
    <div class="parallax-container" id="scene-bottom">
      <h2 data-depth="1">Well, aren't we all?</h2>
    </div>
  </div>
</template>

<script>
import Parallax from "parallax-js";
export default {
  created() {
    var sceneTop = document.getElementById("scene-top");
    var sceneBottom = document.getElementById("scene-bottom");
    var timeoutForAnimation = setTimeout(function() {
      sceneTop = document.getElementById("scene-top");
      sceneBottom = document.getElementById("scene-bottom");
      if (sceneTop && sceneBottom) {
        var parallaxInstance = new Parallax(sceneTop);
        var parallaxInstance = new Parallax(sceneBottom);
        clearInterval(timeoutForAnimation);
      }
    }, 100);
  }
};
</script>

<style>
.error__main {
  text-align: center;
  margin: 0 auto;
  width: 80%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: calc(100vh - 4.1rem);
}

.parallax-container {
  height: 50%;
}

.parallax-container > h2 {
  width: 100%;
}

#scene-top {
  display: flex;
  justify-content: center;
  align-items: flex-end;
}

@media screen and (min-width: 550px) {
  .error__main {
    min-height: calc(100vh - 2rem);
  }
}
</style>
