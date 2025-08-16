// Minimal Prism.js for Kalco Documentation
// This provides basic syntax highlighting for code blocks

(function() {
    'use strict';
    
    // Simple syntax highlighting
    function highlightCode() {
        const codeBlocks = document.querySelectorAll('pre code');
        
        codeBlocks.forEach(block => {
            const text = block.textContent;
            let highlighted = text;
            
            // Basic highlighting patterns
            highlighted = highlighted
                .replace(/(#.*$)/gm, '<span class="token comment">$1</span>')
                .replace(/(\$[^\s]+)/g, '<span class="token keyword">$1</span>')
                .replace(/(\b\w+\.\w+\b)/g, '<span class="token function">$1</span>')
                .replace(/(\b\w+:\s*)/g, '<span class="token property">$1</span>')
                .replace(/(\b\d+\b)/g, '<span class="token number">$1</span>')
                .replace(/(["'][^"']*["'])/g, '<span class="token string">$1</span>');
            
            block.innerHTML = highlighted;
        });
    }
    
    // Initialize when DOM is ready
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', highlightCode);
    } else {
        highlightCode();
    }
    
    // Export for global use
    window.Prism = {
        highlight: highlightCode
    };
})();
