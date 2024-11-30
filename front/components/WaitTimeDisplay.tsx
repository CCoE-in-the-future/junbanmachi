type Props = {
  waitTime: number | null;
};

export default function WaitTimeDisplay({ waitTime }: Props) {
  return (
    <div className="mb-6 text-center">
      <h2 className="text-xl font-medium text-gray-500">
        待ち時間予測:{" "}
        <span className="font-bold text-blue-600">
          {waitTime !== null ? `${waitTime} 分` : "取得中..."}
        </span>
      </h2>
    </div>
  );
}
