<template>
  <div class="main__intro">
    <div class="main__intro-image-container">
      <img class="main__intro-image" src="../../assets/Portrait.webp" alt="Gaurav Saini">
    </div>
    <div class="main__intro-content">
      <div style="margin:auto">
        <vue-typed-js 
          :typeSpeed="125"
          :backSpeed="25"
          :cursorChar="'__'"
          :loop="true"
          :strings="['Gaurav', 'Developer', 'Traveller', 'Dreamer', 'Photographer', 'Foodie', 'Geek']"
        >
          <h1 class="typing"></h1>
        </vue-typed-js>
      </div>
      <div>
        <p>I learn things by doing.</p>
        <div class="main__intro-information">
          <span>
            <a href="tel:+17788583884" style="color:inherit;text-decoration:none;">
              <i class="fa fa-mobile"></i>
              <span>(778) 858-3884</span>
            </a>
          </span>
          <span>
            <i class="fa fa-location-arrow"></i>
            <span>Calgary, AB, Canada</span>
          </span>
        </div>
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
  box-shadow: 0px 0px 4px 4px var(--shadow-color);
  transition: 0.2s all;
}

.main__intro:hover {
  box-shadow: 0px 0px 3px 2px var(--shadow-color);
}

.main__intro-image {
  margin: 2rem;
  margin-bottom: 0.5rem;
  height: 40vmin;
  max-height: 20rem;
  min-height: 10rem;
  border-radius: 50%;
}

.main__intro-content {
  display: flex;
  flex-flow: column;
  text-align: center;
  justify-content: space-around;
  padding: 1rem;
  padding-top: 0;
  flex-grow: 1;
}

.main__intro-content > div {
  height: 5.5rem;
}

.main__intro-content > div > p {
  margin-top: 0;
}

.main__intro-content h1 {
  margin-top: 0;
  font-size: 3rem !important;
}

.main__intro-information {
  display: flex;
  justify-content: center;
  margin: 1rem;
}

.main__intro-information > span * {
  margin: 0.3rem;
  vertical-align: middle;
}

.main__intro-information i {
  color: var(--accent-color);
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
    margin-left: 5rem;
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
