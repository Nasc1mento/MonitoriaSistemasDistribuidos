import express, { Express } from 'express';
import { env } from './config/env';
import { MongoConnection } from './database/mongo.connection';
import ApiRoutes from './routes/routes';


class Server {
    private app: Express;
    private db: MongoConnection
    

    constructor() {
        this.app = express();
        this.db = MongoConnection.getInstance();
    }

    async init(): Promise<void> {
        
        try {
            await this.configDatabase();
            this.configMiddleware();
            this.configRoutes();
            this.startServer();      
        } catch(error) {
            console.error('Erro ao iniciar o servidor', error)
        }
    }

    async configDatabase(): Promise<void> {
        await this.db.connect();
    }

    configRoutes(): void {
        this.app.use('/api', new ApiRoutes().init());
    }

    configMiddleware(): void {
        this.app.use(express.json());
    }

    startServer(): void {
        this.app.listen(env.PORT, () => {
            console.log(`Servidor na porta ${env.PORT}`);
        });
    }
}
new Server().init()