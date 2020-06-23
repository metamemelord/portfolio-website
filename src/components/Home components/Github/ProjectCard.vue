<template>
  <div class="github__card">
    <h3>{{cardData.name|prettifyRepoName}}</h3>
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
    prettifyRepoName(repoName) {
      return repoName.split("-").map(token => token.replace(/\b[a-z]/g, (x) => x.toUpperCase())).join(" ")
    },
    getLastUpdateTime(time) {
      let duration = moment.duration(moment().diff(moment(time)));
      let durationFromMidnight = moment.duration(
        moment().diff(moment(time).startOf("day"))
      );
      time = moment(time);
      if (Math.floor(duration.asYears()) || time.year() !== moment().year()) {
        return `Updated on ${time.format("MMM DD, YYYY")}`;
      } else if (Math.floor(duration.asMonths())) {
        if (time.year() === moment().year() || duration.asMonths() < 3) {
          return `Updated on ${time.format("MMM DD")}`;
        } else {
          return `Updated ${Math.floor(duration.asMonths())} months ago`;
        }
      } else if (Math.floor(duration.asDays())) {
        if (Math.floor(duration.asDays()) > 1)
          return `Updated ${Math.floor(
            durationFromMidnight.asDays()
          )} days ago`;
        return `Updated yesterday`;
      } else if (Math.floor(duration.asHours())) {
        return `Updated ${Math.ceil(duration.asHours())} hours ago`;
      } else if (Math.floor(duration.asMinutes())) {
        if (Math.floor(duration.asMinutes()) > 2)
          return `Updated ${Math.floor(duration.asMinutes())} minutes ago`;
        return `Updated 1 minute ago`;
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
  box-shadow: 0px 0px 3px 1px var(--shadow-color);
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
  color: white;
  background: var(--accent-color);
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
  background: var(--accent-color);
  border-radius: 0px 0px 3px 3px;
  transition: all 0.3s;
  cursor: pointer;
}

.github__card button:hover {
  border: 1px solid var(--accent-color);
  background: transparent;
  color: var(--accent-color);
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