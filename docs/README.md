# Kalco Documentation

This directory contains the documentation website for Kalco, built with the [Just the Docs](https://just-the-docs.github.io/) Jekyll theme.

## Structure

- `index.md` - Main landing page
- `docs/` - Documentation pages organized by topic
- `_config.yml` - Jekyll configuration
- `Gemfile` - Ruby dependencies

## Features

- **Search** - Full-text search across all documentation
- **Navigation** - Hierarchical navigation with sidebar
- **Responsive** - Mobile-friendly design
- **GitHub Pages** - Automatic deployment
- **Callouts** - Highlighted information boxes
- **Code Highlighting** - Syntax highlighting for code blocks

## Local Development

To run the documentation locally:

1. **Install Ruby and Bundler**
   ```bash
   # macOS (with Homebrew)
   brew install ruby
   
   # Ubuntu/Debian
   sudo apt-get install ruby-full build-essential
   
   # Install Bundler
   gem install bundler
   ```

2. **Install Dependencies**
   ```bash
   cd docs
   bundle install
   ```

3. **Run Locally**
   ```bash
   bundle exec jekyll serve
   ```

4. **View Site**
   Open [http://localhost:4000](http://localhost:4000) in your browser

## GitHub Pages

The documentation is automatically deployed to GitHub Pages when pushed to the main branch. The workflow is defined in `.github/workflows/pages.yml`.

## Adding New Pages

1. **Create Markdown File**
   - Use `.md` extension
   - Include front matter with title and navigation order

2. **Front Matter Example**
   ```yaml
   ---
   layout: default
   title: Page Title
   parent: Parent Section
   nav_order: 1
   ---
   ```

3. **Navigation Structure**
   - Use `has_children: true` for parent pages
   - Use `parent:` to specify hierarchy
   - Use `nav_order:` for ordering

## Customization

- **Theme Options** - Modify `_config.yml` under `just_the_docs:`
- **Styling** - Override CSS in `assets/css/`
- **Layouts** - Customize Jekyll layouts in `_layouts/`

## Resources

- [Just the Docs Documentation](https://just-the-docs.github.io/just-the-docs/)
- [Jekyll Documentation](https://jekyllrb.com/)
- [GitHub Pages](https://pages.github.com/)
