<template>
  <main>
    <div class="main__intro">
      <div class="main__intro-image-container">
        <img class="main__intro-image" src="../assets/DP.png" alt="Gaurav Saini">
      </div>
      <div class="main__intro-content">
        <h1>Gaurav Saini</h1>
        <p>I learn things by doing.</p>
        <div class="main__intro-information">
          <span>
            <i class="fa fa-mobile"></i>
            <span>+91-99994-97257</span>
          </span>
          <span>
            <i class="fa fa-location-arrow"></i>
            <span>Bangalore</span>
          </span>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
export default {};
</script>

<style>
main {
  max-width: 70rem;
}
.main__intro {
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: space-between;
  margin: 0.6rem;
  border-radius: 0.3rem;
  box-shadow: 0px 0px 4px 4px rgba(0, 0, 0, 0.15);
  transition: 0.2s all;
}

.main__intro:hover {
  box-shadow: 0px 0px 3px 2px rgba(0, 0, 0, 0.15);
}

.main__intro-image {
  margin: 2rem;
  margin-bottom: 1rem;
  height: 40vmin;
  max-height: 20rem;
  min-height: 10rem;
}

.main__intro-content {
  text-align: center;
  padding: 1rem;
  padding-top: 0;
  flex-grow: 1;
}

.main__intro-content > h1 {
  margin-top: 0;
  font-size: 3rem;
}

.main__intro-information {
  display: flex;
  justify-content: center;
}

.main__intro-information > span * {
  margin: 0.3rem;
  vertical-align: middle;
}

.main__intro-information i {
  color: #008080;
  font-size: 2rem;
}

@media screen and (min-width: 550px) {
  main {
    margin: 0 auto;
  }
  .main__intro {
    flex-direction: row;
    max-width: 70rem;
    margin: 0.6rem;
  }
  .main__intro-image {
    margin: 2rem;
    height: 25vmin;
  }
  /* .main__intro-content {
    text-align: left;
  } */
  .main__intro-content h1 {
    margin-top: inherit;
  }
}

@media screen and (min-width: 71rem) {
  .main__intro {
    margin: 0.6rem auto;
  }
}
</style>

