<script setup lang="ts">
</script>

<template>
  <nav>
    <li>
      <ul><RouterLink to="/" >Home</RouterLink></ul>
    </li>
    <li>
      <ul><RouterLink to="/settings" >Settings</RouterLink></ul>
    </li>
  </nav>
  <RouterView />
</template>
