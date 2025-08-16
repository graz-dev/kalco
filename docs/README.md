# Kalco Documentation

This directory contains the documentation for Kalco, built using the [Just the Docs](https://just-the-docs.github.io/just-the-docs/) Jekyll theme.

## 🏗️ Structure

```
docs/
├── _config.yml              # Jekyll configuration
├── Gemfile                  # Ruby dependencies
├── index.md                 # Home page
├── docs/                    # Documentation pages
│   ├── getting-started/     # Getting started guides
│   │   ├── index.md         # Overview
│   │   ├── installation.md  # Installation guide
│   │   ├── first-run.md     # First run guide
│   │   └── configuration.md # Configuration guide
│   └── commands/            # Command reference
│       ├── index.md         # Commands overview
│       └── export.md        # Export command docs
└── README.md                # This file
```

## 🚀 Local Development

### Prerequisites

- Ruby 3.0+ (recommended: use Homebrew on macOS)
- Bundler gem

### Setup

1. **Install Ruby dependencies:**
   ```bash
   cd docs
   bundle install
   ```

2. **Start local server:**
   ```bash
   bundle exec jekyll serve --host 0.0.0.0 --port 4000
   ```

3. **View site:**
   Open [http://localhost:4000](http://localhost:4000) in your browser

### Build for Production

```bash
bundle exec jekyll build
```

## 🎨 Theme Features

- **Responsive design** - Works on all devices
- **Search functionality** - Full-text search across all content
- **Dark/light mode** - Toggle between color schemes
- **Navigation** - Sidebar navigation with breadcrumbs
- **Callouts** - Highlighted information boxes
- **Git integration** - Last modified timestamps

## 📝 Adding Content

### New Pages

1. Create a new `.md` file in the appropriate directory
2. Add front matter with metadata:
   ```yaml
   ---
   layout: default
   title: Page Title
   nav_order: 1
   parent: Parent Page
   ---
   ```

3. Update navigation by setting `nav_order` values

### Navigation Structure

- Use `nav_order` to control page order
- Use `parent` to create hierarchical navigation
- Use `has_children: true` for pages with sub-pages

## 🔧 Configuration

The `_config.yml` file contains:
- Site metadata (title, description, URL)
- Just the Docs theme configuration
- Jekyll build settings
- Collections and defaults

## 🚀 Deployment

This documentation is automatically deployed to GitHub Pages when changes are pushed to the `main` branch.

## 📚 Resources

- [Just the Docs Documentation](https://just-the-docs.github.io/just-the-docs/)
- [Jekyll Documentation](https://jekyllrb.com/docs/)
- [GitHub Pages](https://pages.github.com/)
