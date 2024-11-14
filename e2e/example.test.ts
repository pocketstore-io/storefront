import { expect, test } from '@nuxt/test-utils/playwright'

test('test', async ({ page, goto }) => {
    await goto('/de', { waitUntil: 'hydration' })
    await expect(page.getByRole("heading", { level: 1 })).toHaveText("Box Office News!");
    await expect(page.getByTestId("p-hero"))
    .toHaveText("Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a id nisi.");
})