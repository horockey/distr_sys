// URL Backend
const backendUrl = '/api';

// Функция для получения IP-адреса клиента
async function getClientIp() {
    try {
        const response = await fetch('https://api.ipify.org?format=json');
        const data = await response.json();
        return data.ip || 'Unknown';
    } catch (error) {
        console.error('Failed to fetch IP address:', error);
        return 'Unknown';
    }
}

// Функция для отправки данных о посещении на Backend
async function logVisit(ip, userAgent) {
    try {
        data = JSON.stringify({ "ip": ip, "userAgent": userAgent })
        const response = await fetch(`${backendUrl}/visit`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: data,
        });
        console.log(data)
        if (!response.ok) {
            throw new Error('Failed to log visit');
        }
        console.log('Visit logged successfully');
    } catch (error) {
        console.log('Error logging visit:' + error);
    }
}

// Функция для получения журнала посещений
async function getVisits() {
    try {
        const response = await fetch(`${backendUrl}/visits`);
        if (!response.ok) {
            throw new Error('Failed to fetch visits');
        }
        const visits = await response.json();
        return visits;
    } catch (error) {
        console.error('Error fetching visits:', error);
        return [];
    }
}

// Функция для отображения журнала посещений в таблице
function renderVisits(visits) {
    const tableBody = document.querySelector('#visit-table tbody');
    tableBody.innerHTML = ''; // Очищаем таблицу

    visits.forEach(visit => {
        const row = document.createElement('tr');

        const timeCell = document.createElement('td');
        timeCell.textContent = new Date(visit.time).toLocaleString();
        row.appendChild(timeCell);

        const ipCell = document.createElement('td');
        ipCell.textContent = visit.ip;
        row.appendChild(ipCell);

        const userAgentCell = document.createElement('td');
        userAgentCell.textContent = visit.userAgent;
        row.appendChild(userAgentCell);

        tableBody.appendChild(row);
    });
}

// Основная функция
async function main() {
    // Получаем IP-адрес и User-Agent
    const ip = await getClientIp();
    const userAgent = navigator.userAgent;

    // Логируем посещение
    await logVisit(ip, userAgent);

    // Получаем и отображаем журнал посещений
    const visits = await getVisits();
    renderVisits(visits);
}

// Запуск основной функции
main();