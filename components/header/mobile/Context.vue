<template>
  <dialog id="modal_mobile_menu" class="modal modal-top px-6">
    <div class="modal-box">
      <HeaderMobileLogo />
      <HeaderMobileSearchBar />
      <HeaderMobileMenu />
      <div class="modal-action">
        <form method="dialog">
          <!-- if there is a button in form, it will close the modal -->
          <button class="btn">{{ $t('general.close') }}</button>
        </form>
      </div>
    </div>
  </dialog>
</template>
