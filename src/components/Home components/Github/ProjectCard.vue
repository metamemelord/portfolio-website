<template>
  <div class="github__card">
    <h3>{{cardData.name.split("-").join(" ")}}</h3>
    <p>{{cardData.description}}</p>
    <div class="github__card-technology-label">
      <span>{{cardData.language}}</span>
    </div>
    <span>{{cardData.updated_at|getLastUpdateTime}}</span>
    <a :href="cardData.html_url" target="_blank" rel="noopener noreferrer">
      <button>Checkout</button>
    </a>
  </div>
</template>

<script>
import * as moment from "moment";
export default {
  props: ["cardData"],
  filters: {
    getLastUpdateTime(time) {
      let duration = moment.duration(moment().diff(moment(time)));
      if (Math.floor(duration.asYears())) {
        return `Updated on ${moment(time).format("MMM DD, YYYY")}`;
      } else if (Math.floor(duration.asMonths())) {
        return `Updated ${Math.ceil(duration.asMonths())} months ago`;
      } else if (Math.floor(duration.asDays())) {
        return `Updated ${Math.ceil(duration.asDays())} days ago`;
      } else if (Math.floor(duration.asHours())) {
        return `Updated ${Math.ceil(duration.asHours())} hours ago`;
      } else if (Math.floor(duration.asMinutes())) {
        return `Updated ${Math.ceil(duration.asMinutes())} minutes ago`;
      } else {
        return "Updated a few moments ago";
      }
    }
  }
};
</script>

<style>
.github__card {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  text-align: center;
  flex-grow: 1;
  flex-basis: 20rem;
  border-radius: 3px;
  margin: 0.5rem;
  box-shadow: 0px 0px 3px 1px rgba(0, 0, 0, 0.15);
}

.github__card h3 {
  margin: 0.9rem 0.7rem;
}

.github__card p {
  margin: 0.9rem;
}

.github__card-technology-label {
  display: flex;
  justify-content: center;
  padding: 0 0.7rem;
  margin: 0.7rem 0;
}

.github__card-technology-label span {
  flex-grow: 1;
  text-align: center;
  padding: 0.5rem;
  margin: 0.5rem;
  background: rgb(255, 161, 38);
  border-radius: 3px;
  max-width: calc(33.33% - 1rem);
}

.github__card > span {
  margin-bottom: 0.9rem;
}

.github__card button {
  bottom: 0px;
  font-size: 0.9rem;
  width: 100%;
  padding: 1rem;
  border: 1px transparent solid;
  color: white;
  background: rgb(255, 161, 38);
  border-radius: 0px 0px 3px 3px;
  transition: all 0.3s;
  cursor: pointer;
}

.github__card button:hover {
  border: 1px solid rgb(255, 161, 38);
  background: transparent;
  color: black;
  border-radius: 0px 0px 3px 3px;
}

@media screen and (min-width: 550px) {
  .github__card {
    max-width: calc(50% - 1rem);
  }
}

@media screen and (min-width: 1000px) {
  .github__card {
    max-width: calc(33.33% - 1rem);
  }
}
</style>