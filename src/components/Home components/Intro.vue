<template>
  <div class="main__intro">
    <div class="main__intro-image-container">
      <img class="main__intro-image" src="../../assets/DP-min.png" alt="Gaurav Saini">
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
</template>

<style>
.main__intro {
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: space-between;
  margin: 1rem;
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
  color: rgb(255, 161, 38);
  font-size: 2rem;
}

@media screen and (min-width: 550px) {
  .main__intro {
    flex-direction: row;
    max-width: 70rem;
    margin: 1rem;
  }
  .main__intro-image {
    margin: 2rem;
    height: 25vmin;
  }
  .main__intro-content h1 {
    margin-top: inherit;
  }
}

@media screen and (min-width: 73rem) {
  .main__intro {
    margin: 1rem auto;
  }
}
</style>
