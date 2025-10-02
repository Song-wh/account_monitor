<script>
  import { onMount, onDestroy } from 'svelte';
  
  export let virtualAccount = "123-456-7890";
  export let authKey = "";
  
  var deposits = [];
  var loading = true;
  var refreshInterval;

  onMount(async () => {
    await loadDeposits();
    
    // 자동 새로고침 설정 (30초마다)
    refreshInterval = setInterval(async () => {
      await loadDeposits();
    }, 30000);
  });

  onDestroy(() => {
    if (refreshInterval) {
      clearInterval(refreshInterval);
    }
  });

  async function loadDeposits() {
    loading = true;
    try {
      const res = await fetch(`/api/deposits?virtual_account=${virtualAccount}`);
      const json = await res.json();
      deposits = json.data || [];
    } catch (error) {
      console.error('Failed to load deposits:', error);
      deposits = [];
    } finally {
      loading = false;
    }
  }

  // 수동 새로고침 함수
  async function manualRefresh() {
    await loadDeposits();
  }

  // Expose refresh function
  export function refresh() {
    loadDeposits();
  }
</script>

<div class="deposit-container">
  <div class="account-info">
    <h2>📊 입금 내역</h2>
    <div class="account-number">
      <span class="label">가상계좌번호:</span>
      <span class="value">{virtualAccount}</span>
      <div class="auto-refresh-indicator">
        <span class="indicator-dot"></span>
        <span class="indicator-text">자동 새로고침 (30초)</span>
      </div>
      <button class="refresh-button" on:click={manualRefresh} disabled={loading}>
        {#if loading}
          <div class="button-spinner"></div>
        {:else}
          🔄
        {/if}
        새로고침
      </button>
    </div>
  </div>

  {#if loading}
    <div class="loading">
      <div class="spinner"></div>
      <p>입금 내역을 불러오는 중...</p>
    </div>
  {:else}
    {#if deposits.length === 0}
      <div class="empty-state">
        <div class="empty-icon">📭</div>
        <h3>입금 내역이 없습니다</h3>
        <p>아직 입금된 내역이 없습니다.</p>
      </div>
    {:else}
      <div class="table-container">
        <table class="deposit-table">
          <thead>
            <tr>
              <th>👤 송금자</th>
              <th>🏦 송금자 계좌번호</th>
              <th>💰 금액</th>
              <th>⏰ 입금시간</th>
            </tr>
          </thead>
          <tbody>
            {#each deposits as d, index}
              <tr class="deposit-row" style="animation-delay: {index * 0.1}s">
                <td class="remitter-name">{d.remitter_name}</td>
                <td class="remitter-account">{d.remitter_account}</td>
                <td class="amount">₩{parseFloat(d.amount).toLocaleString()}</td>
                <td class="created-at">{d.created_at}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
  {/if}
</div>

<style>
  .deposit-container {
    width: 100%;
  }

  .account-info {
    margin-bottom: 30px;
    padding: 20px;
    background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
    border-radius: 12px;
    border-left: 4px solid #667eea;
  }

  .account-info h2 {
    margin: 0 0 15px 0;
    color: #2c3e50;
    font-size: 1.5rem;
    font-weight: 600;
  }

  .account-number {
    display: flex;
    align-items: center;
    gap: 10px;
    flex-wrap: wrap;
  }

  .label {
    color: #6c757d;
    font-weight: 500;
  }

  .value {
    color: #495057;
    font-weight: 600;
    font-family: 'Courier New', monospace;
    background: white;
    padding: 5px 10px;
    border-radius: 6px;
    border: 1px solid #dee2e6;
  }

  .refresh-button {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border: none;
    padding: 8px 16px;
    border-radius: 6px;
    cursor: pointer;
    font-size: 0.9rem;
    font-weight: 500;
    display: flex;
    align-items: center;
    gap: 6px;
    transition: all 0.3s ease;
    margin-left: auto;
  }

  .refresh-button:hover:not(:disabled) {
    transform: translateY(-1px);
    box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
  }

  .refresh-button:disabled {
    opacity: 0.7;
    cursor: not-allowed;
    transform: none;
  }

  .button-spinner {
    width: 14px;
    height: 14px;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-top: 2px solid white;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }

  .auto-refresh-indicator {
    display: flex;
    align-items: center;
    gap: 6px;
    color: #6c757d;
    font-size: 0.8rem;
    margin-left: auto;
    margin-right: 10px;
  }

  .indicator-dot {
    width: 8px;
    height: 8px;
    background: #28a745;
    border-radius: 50%;
    animation: pulse 2s infinite;
  }

  @keyframes pulse {
    0% { opacity: 1; }
    50% { opacity: 0.5; }
    100% { opacity: 1; }
  }

  .loading {
    text-align: center;
    padding: 60px 20px;
  }

  .spinner {
    width: 40px;
    height: 40px;
    border: 4px solid #f3f3f3;
    border-top: 4px solid #667eea;
    border-radius: 50%;
    animation: spin 1s linear infinite;
    margin: 0 auto 20px;
  }

  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }

  .loading p {
    color: #6c757d;
    font-size: 1.1rem;
    margin: 0;
  }

  .empty-state {
    text-align: center;
    padding: 60px 20px;
    color: #6c757d;
  }

  .empty-icon {
    font-size: 4rem;
    margin-bottom: 20px;
  }

  .empty-state h3 {
    margin: 0 0 10px 0;
    color: #495057;
  }

  .empty-state p {
    margin: 0;
    font-size: 1rem;
  }

  .table-container {
    overflow-x: auto;
    border-radius: 12px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  }

  .deposit-table {
    width: 100%;
    border-collapse: collapse;
    background: white;
    border-radius: 12px;
    overflow: hidden;
  }

  .deposit-table thead {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
  }

  .deposit-table th {
    padding: 20px 15px;
    text-align: left;
    font-weight: 600;
    font-size: 0.95rem;
    letter-spacing: 0.5px;
  }

  .deposit-table td {
    padding: 18px 15px;
    border-bottom: 1px solid #f1f3f4;
    vertical-align: middle;
  }

  .deposit-row {
    transition: all 0.3s ease;
    animation: slideIn 0.5s ease forwards;
    opacity: 0;
    transform: translateY(20px);
  }

  @keyframes slideIn {
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  .deposit-row:hover {
    background: linear-gradient(135deg, #f8f9ff 0%, #f0f2ff 100%);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(102, 126, 234, 0.15);
  }

  .deposit-row:last-child td {
    border-bottom: none;
  }

  .remitter-name {
    font-weight: 600;
    color: #2c3e50;
  }

  .remitter-account {
    font-family: 'Courier New', monospace;
    color: #495057;
    font-size: 0.9rem;
  }

  .amount {
    font-weight: 700;
    color: #27ae60;
    font-size: 1.1rem;
  }

  .created-at {
    color: #6c757d;
    font-size: 0.9rem;
  }

  @media (max-width: 768px) {
    .deposit-table th,
    .deposit-table td {
      padding: 12px 8px;
      font-size: 0.85rem;
    }

    .account-info {
      padding: 15px;
    }

    .account-number {
      flex-direction: column;
      align-items: flex-start;
      gap: 5px;
    }
  }
</style>
