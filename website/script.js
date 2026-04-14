const body = document.body;
const owner = body.dataset.owner || "AasishDairelSahayaGrinspan";
const repo = body.dataset.repo || "repoant";

const repoUrl = `https://github.com/${owner}/${repo}`;
const apiUrl = `https://api.github.com/repos/${owner}/${repo}`;

const starCountEl = document.getElementById("star-count");
const starLinkEl = document.getElementById("star-link");
const repoLinkEl = document.getElementById("repo-link");

if (starLinkEl) {
  starLinkEl.href = `${repoUrl}/stargazers`;
}

if (repoLinkEl) {
  repoLinkEl.href = repoUrl;
}

function formatStars(value) {
  return new Intl.NumberFormat("en", { notation: "compact" }).format(value);
}

function animateCount(endValue) {
  if (!starCountEl) {
    return;
  }

  const duration = 700;
  const startTime = performance.now();

  function tick(now) {
    const progress = Math.min((now - startTime) / duration, 1);
    const eased = 1 - Math.pow(1 - progress, 3);
    const current = Math.round(endValue * eased);
    starCountEl.textContent = formatStars(current);

    if (progress < 1) {
      requestAnimationFrame(tick);
    }
  }

  requestAnimationFrame(tick);
}

if (starCountEl) {
  fetch(apiUrl)
    .then((response) => {
      if (!response.ok) {
        throw new Error("Could not fetch repo data");
      }
      return response.json();
    })
    .then((data) => {
      const stars = Number(data.stargazers_count || 0);
      animateCount(stars);
    })
    .catch(() => {
      // Safe fallback if GitHub API is unavailable.
      starCountEl.textContent = "N/A";
    });
}
