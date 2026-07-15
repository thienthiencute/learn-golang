/**
 * Sidebar Navigation Handler
 * Xử lý active state cho sidebar menu với support cho templates.SafeURL()
 */

class SidebarNavigation {
    constructor() {
        this.currentPath = window.location.pathname;
        this.currentOrigin = window.location.origin;
        this.sidebarElement = document.getElementById("navbar-nav");
        this.twoColumnElement = document.getElementById("two-column-menu");
        
        this.init();
    }

    init() {
        // Chờ DOM load xong
        if (document.readyState === 'loading') {
            document.addEventListener('DOMContentLoaded', () => {
                this.initMenu();
            });
        } else {
            this.initMenu();
        }
    }

    initMenu() {
        // Clear all active states first
        this.clearActiveStates();

        // Always rely on URL-based detection
        this.setActiveMenu();
    }

    /**
     * Normalize URL để so sánh
     * Xử lý cả relative path và absolute URL từ templates.SafeURL()
     */
    normalizeURL(url) {
        if (!url) return '';
        
        // Nếu là absolute URL, lấy pathname
        if (url.startsWith('http://') || url.startsWith('https://')) {
            try {
                const urlObj = new URL(url);
                return urlObj.pathname;
            } catch (e) {
                return url;
            }
        }
        
        // Nếu là relative path, return as is
        return url;
    }

    /**
     * Tìm link active dựa trên current path
     */
    findActiveLink() {
        if (!this.sidebarElement) return null;

        const allLinks = this.sidebarElement.querySelectorAll('a[href]');
        let exactMatch = null;
        let bestPartialMatch = null;
        let longestMatchLength = 0;

        for (const link of allLinks) {
            const href = link.getAttribute('href');
            const normalizedHref = this.normalizeURL(href);
            
            // Bỏ qua các link có data-bs-toggle (collapse toggles)
            if (link.getAttribute('data-bs-toggle')) {
                continue;
            }

            // Bỏ qua link rỗng
            if (!normalizedHref || normalizedHref === '#') {
                continue;
            }

            // Exact match
            if (normalizedHref === this.currentPath) {
                exactMatch = link;
                break;
            }

            // Partial match (currentPath starts with href)
            if (normalizedHref !== '/' && this.currentPath.startsWith(normalizedHref)) {
                if (normalizedHref.length > longestMatchLength) {
                    longestMatchLength = normalizedHref.length;
                    bestPartialMatch = link;
                }
            }

            // Special case for home page
            if ((normalizedHref === '' || normalizedHref === '/') && 
                (this.currentPath === '/' || this.currentPath === '')) {
                exactMatch = link;
                break;
            }
        }

        return exactMatch || bestPartialMatch;
    }

    /**
     * Xóa tất cả active states
     */
    clearActiveStates() {
        if (!this.sidebarElement) return;

        // Xóa active class từ tất cả nav-links
        const allNavLinks = this.sidebarElement.querySelectorAll('.nav-link');
        allNavLinks.forEach(link => {
            link.classList.remove('active');
        });

        // Đóng tất cả collapse menus
        const allCollapses = this.sidebarElement.querySelectorAll('.collapse');
        allCollapses.forEach(collapse => {
            collapse.classList.remove('show');
        });

        // Reset aria-expanded
        const allToggleLinks = this.sidebarElement.querySelectorAll('[data-bs-toggle="collapse"]');
        allToggleLinks.forEach(link => {
            link.setAttribute('aria-expanded', 'false');
            link.classList.remove('active');
        });

        // Clear two-column specific states
        if (this.twoColumnElement) {
            const allTwoColumnLinks = this.twoColumnElement.querySelectorAll('a');
            allTwoColumnLinks.forEach(link => {
                link.classList.remove('active');
            });
        }

        // Remove twocolumn-item-show class
        const allTwoColumnItems = document.querySelectorAll('.twocolumn-item-show');
        allTwoColumnItems.forEach(item => {
            item.classList.remove('twocolumn-item-show');
        });
    }

    /**
     * Set active state cho link và parent menus
     */
    setActiveState(activeLink) {
        if (!activeLink) return;

        // Set active cho link hiện tại
        activeLink.classList.add('active');

        // Xử lý parent collapse menus
        this.handleParentCollapses(activeLink);

        // Xử lý two-column layout nếu cần
        this.handleTwoColumnLayout(activeLink);

        // Scroll to active menu item
        this.scrollToActiveItem(activeLink);
    }

    /**
     * Xử lý parent collapse menus
     */
    handleParentCollapses(activeLink) {
        let currentElement = activeLink;

        while (currentElement) {
            const parentCollapse = currentElement.closest('.collapse.menu-dropdown');
            if (!parentCollapse) break;

            // Show parent collapse
            parentCollapse.classList.add('show');

            // Set active cho toggle button
            const toggleButton = parentCollapse.parentElement.querySelector('[data-bs-toggle="collapse"]');
            if (toggleButton) {
                toggleButton.classList.add('active');
                toggleButton.setAttribute('aria-expanded', 'true');
            }

            // Move to next level
            currentElement = parentCollapse.parentElement;
        }
    }

    /**
     * Xử lý two-column layout
     */
    handleTwoColumnLayout(activeLink) {
        const isTwoColumn = document.documentElement.getAttribute('data-layout') === 'twocolumn';
        if (!isTwoColumn || !this.twoColumnElement) return;

        // Set active cho two-column menu
        const href = activeLink.getAttribute('href');
        const normalizedHref = this.normalizeURL(href);
        
        const twoColumnLink = this.twoColumnElement.querySelector(`[href="${href}"], [href="${normalizedHref}"]`);
        if (twoColumnLink) {
            twoColumnLink.classList.add('active');
        }

        // Xử lý parent collapse cho two-column
        const parentCollapse = activeLink.closest('.collapse.menu-dropdown');
        if (parentCollapse) {
            const parentElement = parentCollapse.parentElement;
            
            // Check for nested collapse
            const grandParentCollapse = parentElement.closest('.collapse.menu-dropdown');
            if (grandParentCollapse) {
                // Nested menu case
                grandParentCollapse.parentElement.classList.add('twocolumn-item-show');
                
                const menuId = grandParentCollapse.getAttribute('id');
                if (menuId) {
                    const twoColumnToggle = this.twoColumnElement.querySelector(`[href="#${menuId}"]`);
                    if (twoColumnToggle) {
                        twoColumnToggle.classList.add('active');
                    }
                }
            } else {
                // Single level menu
                parentElement.classList.add('twocolumn-item-show');
                const menuId = parentCollapse.getAttribute('id');
                if (menuId) {
                    const twoColumnToggle = this.twoColumnElement.querySelector(`[href="#${menuId}"]`);
                    if (twoColumnToggle) {
                        twoColumnToggle.classList.add('active');
                    }
                }
            }
        }
    }

    /**
     * Scroll to active menu item
     */
    scrollToActiveItem(activeLink) {
        setTimeout(() => {
            const offset = activeLink.offsetTop;
            if (offset > 300) {
                const verticalMenu = document.querySelector('.app-menu');
                if (verticalMenu && verticalMenu.querySelector('.simplebar-content-wrapper')) {
                    setTimeout(() => {
                        const scrollTop = offset === 330 ? offset + 85 : offset;
                        verticalMenu.querySelector('.simplebar-content-wrapper').scrollTop = scrollTop;
                    }, 0);
                }
            }
        }, 250);
    }

    /**
     * Main function để set active menu
     */
    setActiveMenu() {
        // Clear all active states first
        this.clearActiveStates();

        // Find and set active link
        const activeLink = this.findActiveLink();
        if (activeLink) {
            this.setActiveState(activeLink);
            // No longer store active link in local storage
        } else {
            // Fallback: nếu không tìm thấy active link, show two-column panel
            const isTwoColumn = document.documentElement.getAttribute('data-layout') === 'twocolumn';
            if (isTwoColumn) {
                document.body.classList.add('twocolumn-panel');
            }
        }
    }

    /**
     * Update active menu khi URL thay đổi (cho SPA)
     */
    updateActiveMenu() {
        this.currentPath = window.location.pathname;
        this.setActiveMenu();
    }
}

// Export cho sử dụng global
window.SidebarNavigation = SidebarNavigation;

// Auto initialize
let sidebarNav = null;

// Initialize when DOM is ready
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', () => {
        sidebarNav = new SidebarNavigation();
    });
} else {
    sidebarNav = new SidebarNavigation();
}

// Export instance
export default sidebarNav;
