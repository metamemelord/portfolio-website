<template>
  <div class="main__intro">
    <div class="main__intro-image-container">
      <img class="main__intro-image" src="../../assets/Portrait.webp" alt="Gaurav Saini">
    </div>
    <div class="main__intro-content">
      <div class="main__typing" style="margin:0.5rem auto 0 auto">
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
            <a href="tel:+17788583884">
              <i class="fa fa-phone"></i>
              <span>(778) 858-3884</span>
            </a>
          </span>
          <span>
            <i class="fa fa-location-arrow"></i>
            <span>Calgary, Canada</span>
          </span>
        </div>
        <div class="main__intro-information">
          <span>
            <a href="https://instagram.gaurav.dev" target="_blank" class="icon-ig">
              <i class="fa-brands fa-instagram"></i>
            </a>
          </span>
          <span>
            <a href="https://linkedin.gaurav.dev" target="_blank" class="icon-in">
              <i class="fa-brands fa-linkedin"></i>
            </a>
          </span>
          <span>
            <a href="https://api.whatsapp.com/send?phone=919999497257" class="icon-wa" target="_blank">
              <i class="fa-brands fa-whatsapp"></i>
            </a>
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
  flex-grow: 1;
}

.main__typing {
  min-height: 6rem;
}

.main__typing h1 {
  margin-top: 0;
  font-size: 3rem !important;
}

.main__intro-content > div > p {
  margin-top: 0;
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

.main__intro-information a {
  color: inherit;
  text-decoration: none;
}

.main__intro-information i {
  color: var(--accent-color);
  font-size: 2rem;
  transition: 0.3s;
}

.main__intro-information .icon-ig i:hover,
.main__intro-information .icon-ig i:focus {
  color: #833AB4;
}

.main__intro-information .icon-in i:hover,
.main__intro-information .icon-in i:focus {
  color: #0A66C2;
}

.main__intro-information .icon-wa i:hover,
.main__intro-information .icon-wa i:focus {
  color: #28D146;
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
