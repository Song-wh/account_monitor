<script>
  import Deposit from './routes/Deposit.svelte';
  import LoginForm from './components/LoginForm.svelte';
  
  let isAuthenticated = false;
  let virtualAccount = '';
  let authKey = '';

  function handleLoginSuccess(event) {
    virtualAccount = event.detail.virtualAccount;
    authKey = event.detail.authKey;
    isAuthenticated = true;
  }

  function handleLogout() {
    isAuthenticated = false;
    virtualAccount = '';
    authKey = '';
  }
</script>

<main>
  <div class="container">
    <header class="header">
      <h1>🏦 가상계좌 입금 모니터링 시스템</h1>
      <p class="subtitle">실시간 입금 내역을 확인하세요</p>
      
      {#if isAuthenticated}
        <div class="user-info">
          <span class="account-display">계좌: {virtualAccount}</span>
          <button class="logout-button" on:click={handleLogout}>
            🚪 로그아웃
          </button>
        </div>
      {/if}
    </header>
    
    <div class="content">
      {#if isAuthenticated}
        <Deposit {virtualAccount} {authKey} />
      {:else}
        <LoginForm on:login-success={handleLoginSuccess} />
      {/if}
    </div>
  </div>
</main>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
  }

  main {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 20px;
  }

  .container {
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    border-radius: 20px;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
    padding: 40px;
    max-width: 1200px;
    width: 100%;
    box-sizing: border-box;
  }

  .header {
    text-align: center;
    margin-bottom: 40px;
  }

  h1 {
    color: #2c3e50;
    font-size: 2.5rem;
    font-weight: 700;
    margin: 0 0 10px 0;
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .subtitle {
    color: #7f8c8d;
    font-size: 1.1rem;
    margin: 0 0 20px 0;
    font-weight: 400;
  }

  .user-info {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 20px;
    margin-top: 10px;
  }

  .account-display {
    background: rgba(255, 255, 255, 0.2);
    padding: 8px 16px;
    border-radius: 20px;
    color: white;
    font-weight: 500;
    font-size: 0.9rem;
  }

  .logout-button {
    background: rgba(255, 255, 255, 0.2);
    border: 1px solid rgba(255, 255, 255, 0.3);
    color: white;
    padding: 8px 16px;
    border-radius: 20px;
    cursor: pointer;
    font-size: 0.9rem;
    transition: all 0.3s ease;
  }

  .logout-button:hover {
    background: rgba(255, 255, 255, 0.3);
    transform: translateY(-1px);
  }

  .content {
    background: white;
    border-radius: 15px;
    padding: 30px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
  }

  @media (max-width: 768px) {
    .container {
      padding: 20px;
      margin: 10px;
    }
    
    h1 {
      font-size: 2rem;
    }
    
    .content {
      padding: 20px;
    }
  }
</style>
