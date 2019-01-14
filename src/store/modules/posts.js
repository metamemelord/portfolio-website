import blogPosts from "../../data/BlogPosts";

const state = {
  blogPosts: []
};

const mutations = {
  SET_POSTS(state, blogPosts) {
    state.blogPosts = blogPosts;
  }
};

const actions = {
  initPosts: ({ commit }) => {
    commit("SET_POSTS", blogPosts);
  }
};

const getters = {
  blogPosts: state => {
    return state.blogPosts;
  },
  blogPostById: state => id => {
    return state.blogPosts.filter(post => {
      if (post.id === id) return post;
    })[0];
  }
};

export default {
  state,
  mutations,
  actions,
  getters
};
