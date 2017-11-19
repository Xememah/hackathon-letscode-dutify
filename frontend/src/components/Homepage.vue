
<template>
</template>
<script>
export default {
  data() {
    return {
      dialog: false,
      progress: [
        {
          points: 72,
          title: "Weronika",
          avatar: "/static/img/avatars/1.jpg",
          disabled: true
        },
        {
          points: 51,
          title: "Marek",
          avatar: "/static/img/avatars/2.jpg",
          disabled: false        
        },
        {
          points: 47,
          title: "Gracjan",
          avatar: "/static/img/avatars/3.jpg ",
          disabled: true
        }
      ],
      actions: [
        {
          title: "Bathroom cleaning",
          points: 10,
          icon: "home"
        },
        {
          title: "Cooking dinner",
          points: 7,
          icon: "restaurant"
        },
        {
          title: "Dishwashing",
          points: 3,
          icon: "local_dining"
        },
        {
          title: "Shopping for groceries",
          points: 2,
          icon: "shopping_basket"
        }
      ]
    };
  }
};
</script>

<style>
  .disabled {
    opacity: 0.5;
  }
</style>