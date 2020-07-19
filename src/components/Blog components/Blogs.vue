<template>
  <main>
    <div class="blogs__tags-filter" v-if="getTagsByCount().length">
      <b>Top tags:</b>
      <div>
        <span
          v-for="(tag,idx) in getTagsByCount()"
          :key="idx"
          :value="tag[1].toLowerCase()"
          @click="toggleTagInData($event)"
        >{{tag[1]|capitalize}}</span>
      </div>
    </div>
    <blog-post v-for="(blogPost, idx) in blogPosts" :key="idx" :blogPost="blogPost"></blog-post>
  </main>
</template>

<script>
import blogPost from "./Blog.vue";

export default {
  data() {
    return {
      tags: []
    };
  },
  components: {
    blogPost
  },
  methods: {
    toggleTagInData(event) {
      let el = event.target;
      let tag = el.getAttribute("value");
      let idx = this.tags.indexOf(tag);
      if (idx == -1) {
        this.tags.push(tag);
        el.classList.add("active-tag");
      } else {
        this.tags.splice(idx, 1);
        el.className = "";
        el.classList.remove("active-tag");
      }
    },
    getTagsByCount() {
      let tags = {};
      for (let post of this.$store.getters.blogPosts) {
        for (let tag of post.tags) {
          if (tag in tags) {
            tags[tag]++;
          } else {
            tags[tag] = 1;
          }
        }
      }
      let res = [];
      for (let tag in tags) {
        res.push([tags[tag], tag]);
      }
      return res
        .sort((el1, el2) => parseInt(el1[0]) < parseInt(el2[0]))
        .slice(0, 7);
    }
  },
  computed: {
    blogPosts() {
      if (this.tags.length == 0) {
        return this.$store.getters.blogPosts;
      }
      return this.$store.getters.blogPosts.filter(el => {
        for (let tag of el.tags) {
          if (this.tags.includes(tag.toLowerCase())) {
            return true;
          }
        }
        return false;
      });
    }
  },
  beforeMount() {
    this.$http
      .get("api/wordpress")
      .then(res => {
        if (res.status == 200) this.$store.dispatch("setPosts", res.body);
        else throw "Wordpress call failed";
      })
      .catch(() => this.$store.dispatch("initPosts"));
  }
};
</script>
<style>
main {
  max-width: 72rem;
  margin: auto;
  margin-top: 7.3rem;
  min-height: calc(100vh - 10.8rem);
}
.blogs__tags-filter {
  display: flex;
  flex-wrap: wrap;
  align-content: space-between;
  justify-content: center;
  position: relative;
  top: -1.4rem;
  padding: 1.25rem;
  padding-left: 2rem;
  margin: 1rem;
  border-radius: 0 0 0.3rem 0.3rem;
  box-shadow: 0px 0px 7px 4px var(--shadow-color);
}

.blogs__tags-filter div {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.blogs__tags-filter span {
  display: inline-block;
  margin: 0.3rem;
  padding: 0.3rem;
  background: var(--accent-color);
  color: #2e3342;
  border-radius: 0.2rem;
  cursor: pointer;
  transition: background, color 0.2s linear;
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

@media screen and (min-width: 550px) {
  main {
    min-height: calc(100vh - 9rem);
  }

  .blogs__tags-filter {
    display: block;
  }

  .blogs__tags-filter div {
    display: inline;
  }

  .blogs__tags-filter span {
    margin: 0.5rem;
    padding: 0.5rem;
  }
}

.active-tag {
  background: rgb(204, 128, 28) !important;
  color: white !important;
  box-shadow: 0 0 0.4rem var(--shadow-color);
}
</style>
