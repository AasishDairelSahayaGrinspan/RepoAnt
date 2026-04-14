# RepoAnt Website Plan and Documentation

## Scope

This document covers the marketing website in `website/`, including scenario simulation media, content pages, and GitHub Pages deployment.

## Delivery Plan

1. Build a scenario-first showcase section on the homepage.
2. Add SEO pages: features deep dive and documentation.
3. Keep sitemap and robots aligned with published URLs.
4. Deploy from `website/` using GitHub Pages workflow.
5. Maintain media and docs together after each product change.

## Implemented Pages

- `website/index.html`: main landing page with star counter and simulated scenario gallery.
- `website/features.html`: SEO-focused deep-dive content.
- `website/documentation.html`: full operational and user documentation.

## Scenario Simulation Assets

- `website/assets/scenario-login.svg`
- `website/assets/scenario-list.svg`
- `website/assets/scenario-single-delete.svg`
- `website/assets/repoant-screenshot.svg` (multi-delete)
- `website/assets/scenario-protected.svg`
- `website/assets/scenario-token-missing.svg`
- `website/assets/repoant-demo.gif` (replace with real GIF recording)

## Media Update Process

1. Record a terminal session that includes login, list, single delete, multi delete, and protected repo behavior.
2. Export a compact GIF for `website/assets/repoant-demo.gif`.
3. Capture still screenshots (or regenerate SVG mocks) for each scenario state.
4. Keep captions and alt text aligned with actual behavior.

## GitHub Pages Deployment

Workflow file: `.github/workflows/pages.yml`

Behavior:
- Runs on push to `main`.
- Uploads `website/` as the Pages artifact.
- Deploys to project Pages.

Published URLs:
- `https://aasishdairelsahayagrinspan.github.io/RepoAnt/`
- `https://aasishdairelsahayagrinspan.github.io/RepoAnt/features.html`
- `https://aasishdairelsahayagrinspan.github.io/RepoAnt/documentation.html`

## Operations Checklist

Before pushing:
- Confirm links and canonical URLs use `/RepoAnt/` path casing.
- Run local preview from `website/`.
- Ensure `website/sitemap.xml` includes all public pages.
- Confirm `website/robots.txt` points to the right sitemap URL.

After pushing:
- Verify Pages workflow succeeds in Actions.
- Validate each public URL in a browser.
- Confirm star counter still renders.

## Troubleshooting

### 404 on Pages
- Use `https://aasishdairelsahayagrinspan.github.io/RepoAnt/` with exact case.
- Confirm Pages source is GitHub Actions in repository settings.
- Check workflow permissions include `pages: write` and `id-token: write`.

### Media Missing
- Ensure referenced files exist in `website/assets/`.
- Check spelling in `index.html` `src` attributes.

### Star Count Not Loading
- Check GitHub API availability or request limits.
- Verify repository owner/repo dataset values in page body.
