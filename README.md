<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a id="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/Auntjoestar/cs50NameChanger">
    <img src="./resources/appicon.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">CS50 Name Changer</h3>

  <p align="center">
    An app created to change the filenames of images withe the oficial formate required at CS50x Nicaragua.
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](./resources/Screenshot%202024-12-06%20225255.png)

This app starts as a personal projecto to automate the name changing of the pictures for [CS50 y Web50 Nicaragua](https://code-fu.net.ni/)'s SmugMug.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Wails][Wails]][Wails-url]
* [![Vue][Vue.js]][Vue-url]
* [![Sass][Sass]][Sass-url]
* [![Bootstrap][Bootstrap.com]][Bootstrap-url]
* [![Sqlite][Sqlite.com]][sqlite-url]
* [![Go][Go.com]][Go-url]
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple example steps.

### Prerequisites
* Go
  ```sh
  sudo apt install go
  ```
* npm
  ```sh
  npm install npm@latest -g
  ```
* Wails
  ```sh
  go install github.com/wailsapp/wails/v2/cmd/wails@latest
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/auntjoestar/cs50NameChanger.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
5. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin auntjoestar/cs50NameChanger
   git remote -v # confirm the changes
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

### 1. Installation  
Download the application from the website in your preferred version. Upon installation, a database with preloaded content (sufficient for 2025 photos) will be created automatically.

### 2. Initial Connection  
When you open the app, you'll see a screen to connect to the database. Click the indicated button to load a form where you can select the **program**, **cycle**, **week**, and **group**.

### 3. File Selection  
- Use the folder icon button to select files for renaming.  
- A table will display the selected files, allowing you to deselect them if needed.  
- Add more files or hide the table using the controls at the top.

### 4. Configure the Index  
Choose the starting index for renaming. By default, the app will use the last recorded index for the matching pattern in the target folder. Indices lower than the last recorded value cannot be selected.

### 5. Rename Files  
Click the **"Rename"** button to rename the files according to the desired format. The program will create a folder (if it doesn’t already exist) within the directory of the executable.

### 6. Admin Access  
If the required **program/cycle/week/group** isn’t available, use the admin button in the bottom-right corner (database icon). This allows you to view, modify, delete, or add new data as needed.



<p align="right">(<a href="#readme-top">back to top</a>)</p>




<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Top contributors:

<a href="https://github.com/auntjoestar/cs50NameChanger/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=auntjoestar/cs50NameChanger" alt="contrib.rocks image" />
</a>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Auntjoestar - sebalsor23@gmail.com

Project Link: [https://github.com/auntjoestar/cs50NameChanger](https://github.com/auntjoestar/cs50NameChanger)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* []()
* []()
* []()

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/auntjoestar/cs50NameChanger.svg?style=for-the-badge
[contributors-url]: https://github.com/auntjoestar/cs50NameChanger/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/auntjoestar/cs50NameChanger.svg?style=for-the-badge
[forks-url]: https://github.com/auntjoestar/cs50NameChanger/network/members
[stars-shield]: https://img.shields.io/github/stars/auntjoestar/cs50NameChanger.svg?style=for-the-badge
[stars-url]: https://github.com/auntjoestar/cs50NameChanger/stargazers
[issues-shield]: https://img.shields.io/github/issues/auntjoestar/cs50NameChanger.svg?style=for-the-badge
[issues-url]: https://github.com/auntjoestar/cs50NameChanger/issues
[license-shield]: https://img.shields.io/github/license/auntjoestar/cs50NameChanger.svg?style=for-the-badge
[license-url]: https://github.com/auntjoestar/cs50NameChanger/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[product-screenshot]: images/screenshot.png
[Wails]: https://img.shields.io/badge/Wails-DF0000?style=for-the-badge&logo=wails&logoColor=white
[wails-url]: https://wails.io/
[Sass]: https://img.shields.io/badge/Sass-CC6699?style=for-the-badge&logo=sass&logoColor=white
[Sass-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[Sqlite.com]: https://img.shields.io/badge/Sqlite-003B57?style=for-the-badge&logo=sqlite&logoColor=white
[Sqlite-url]: https://www.sqlite.org/
[Go.com]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
