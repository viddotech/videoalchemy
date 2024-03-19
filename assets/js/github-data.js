document.addEventListener("DOMContentLoaded", function() {
    // Replace 'your-username/your-repository' with your actual GitHub repo path
    const repo = 'viddotech/videoalchemy';

    fetch(`https://api.github.com/repos/${repo}`)
        .then(response => response.json())
        .then(data => {
            const stars = data.stargazers_count;
            const header = document.querySelector('.md-header-nav__title');
            const starsElement = document.createElement('span');
            starsElement.innerHTML = `⭐ ${stars}`;
            starsElement.style.marginLeft = '10px';
            if (header) {
                header.appendChild(starsElement);
            }
        })
        .catch(error => console.error('Error fetching GitHub stars:', error));
});
